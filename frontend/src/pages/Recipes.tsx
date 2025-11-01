import { useEffect, useState, useCallback } from "react";
import { RecipeCard } from "@/components/recipes/RecipeCard";
import { RecipeDrawer } from "@/components/recipes/RecipeDrawer";
import { Loader2, Search, Pencil, Trash2, Plus } from "lucide-react";
import { Button } from "@/components/ui/button";
import debounce from "lodash.debounce";
import {
  InputGroup,
  InputGroupAddon,
  InputGroupInput,
} from "../components/ui/input-group";
import type {
  GetRecipesResponseDto,
  RecipeDto,
  RecipeInput,
} from "../types/recipe";
import { appConfig } from "../lib/config";
import { useAuth } from "../context/AuthContext";
import { Topbar } from "../components/layout/Topbar";
import { RecipeForm } from "@/components/recipes/RecipeForm";
import { ConfirmDialog } from "../components/layout/ConfirmDialog";

const LIMIT_OPTIONS = [5, 10, 20] as const;

const Recipes = () => {
  const { token, logout, user, isAdmin } = useAuth();

  const [recipes, setRecipes] = useState<RecipeDto[]>([]);
  const [loading, setLoading] = useState(true);

  const [page, setPage] = useState(1);
  const [limit, setLimit] = useState(10);
  const [totalPages, setTotalPages] = useState(1);

  const [search, setSearch] = useState("");
  const [selectedRecipe, setSelectedRecipe] = useState<RecipeDto | null>(null);

  const [showForm, setShowForm] = useState(false);
  const [editingRecipe, setEditingRecipe] = useState<RecipeDto | null>(null);
  const [deletingRecipe, setDeletingRecipe] = useState<RecipeDto | null>(null);

  const updateQueryParam = (key: string, value: string | null) => {
    const params = new URLSearchParams(window.location.search);

    if (value === null || value === "") {
      params.delete(key);
    } else {
      params.set(key, value);
    }

    const queryString = params.toString();
    const newUrl = queryString
      ? `${window.location.pathname}?${queryString}`
      : window.location.pathname;

    window.history.replaceState({}, "", newUrl);
  };

  const fetchRecipes = async (page: number, limit: number, search: string) => {
    try {
      setLoading(true);

      const query = new URLSearchParams({
        page: String(page),
        limit: String(limit),
        ...(search && { name: search }),
      });

      const response = await fetch(
        `${appConfig.baseApiUrl}/recipes?${query.toString()}`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      const data: GetRecipesResponseDto = await response.json();

      setRecipes(data.recipes);
      setTotalPages(data.meta.totalPages);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const debouncedFetch = useCallback(
    debounce((page, limit, search) => {
      fetchRecipes(page, limit, search);
    }, 200),
    []
  );

  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const editId = params.get("edit");

    if (!editId) return;

    console.log(recipes);

    if (recipes.length === 0) return;

    const found = recipes.find((recipe) => recipe.id === editId);
    if (!found) return;

    setEditingRecipe(found);
    setShowForm(true);
  }, [recipes]);

  useEffect(() => {
    debouncedFetch(page, limit, search);
  }, [page, limit, search]);

  const handleSubmit = async (data: RecipeInput) => {
    const url = editingRecipe
      ? `${appConfig.baseApiUrl}/recipes/${editingRecipe.id}`
      : `${appConfig.baseApiUrl}/recipes`;
    const method = editingRecipe ? "PUT" : "POST";

    await fetch(url, {
      method,
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(data),
    });

    closeForm();
    debouncedFetch(page, limit, search);
  };

  const closeForm = () => {
    setShowForm(false);
    setEditingRecipe(null);

    updateQueryParam("edit", null);
  };

  const openEditForm = (recipe: RecipeDto) => {
    const params = new URLSearchParams(window.location.search);
    params.set("edit", recipe.id.toString());
    window.history.replaceState(
      {},
      "",
      `${window.location.pathname}?${params}`
    );

    setEditingRecipe(recipe);
    setShowForm(true);
  };

  return (
    <div className="min-h-screen px-6 py-6 bg-gray-50 dark:bg-gray-900 text-gray-800 dark:text-gray-200">
      <Topbar user={user!} onLogout={logout} />

      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-6 gap-4">
        <div className="w-[17%] min-w-[180px]">
          <InputGroup>
            <InputGroupInput
              placeholder="Search recipes..."
              value={search}
              onChange={(event) => {
                setPage(1);
                setSearch(event.target.value);
              }}
            />
            <InputGroupAddon>
              <Search />
            </InputGroupAddon>
          </InputGroup>
        </div>

        {isAdmin && (
          <Button
            onClick={() => {
              updateQueryParam("edit", null);
              setShowForm(true);
            }}
            className="flex items-center gap-2"
          >
            <Plus className="h-4 w-4" /> Add a recipe
          </Button>
        )}

        <div className="flex items-center gap-4 ml-auto">
          <div className="flex items-center gap-2">
            <span className="text-sm">Recipes per page:</span>
            <select
              value={limit}
              onChange={(e) => {
                setLimit(Number(e.target.value));
                setPage(1);
              }}
              className="px-2 py-1 rounded bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-700 text-sm"
            >
              {LIMIT_OPTIONS.map((n) => (
                <option key={n} value={n}>
                  {n}
                </option>
              ))}
            </select>
          </div>

          <div className="flex gap-2">
            <Button
              variant="outline"
              size="sm"
              disabled={page === 1}
              onClick={() => setPage((p) => Math.max(1, p - 1))}
            >
              Previous
            </Button>

            <Button
              variant="outline"
              size="sm"
              disabled={page === totalPages}
              onClick={() => setPage((p) => p + 1)}
            >
              Next
            </Button>
          </div>
        </div>
      </div>

      {loading ? (
        <div className="w-full flex justify-center py-20">
          <Loader2 className="w-10 h-10 animate-spin text-gray-600 dark:text-gray-300" />
        </div>
      ) : (
        <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-6">
          {recipes.map((recipe) => (
            <div key={recipe.id} className="relative">
              <RecipeCard
                recipe={recipe}
                onDetails={() => setSelectedRecipe(recipe)}
              />
              {isAdmin && (
                <div className="absolute top-3 right-3 flex gap-2">
                  <Button
                    size="icon"
                    variant="ghost"
                    onClick={() => openEditForm(recipe)}
                  >
                    <Pencil className="h-4 w-4" />
                  </Button>
                  <Button
                    size="icon"
                    variant="ghost"
                    onClick={() => setDeletingRecipe(recipe)}
                  >
                    <Trash2 className="h-4 w-4 text-red-500" />
                  </Button>
                </div>
              )}
            </div>
          ))}
        </div>
      )}

      {selectedRecipe && (
        <RecipeDrawer
          recipe={selectedRecipe}
          onClose={() => setSelectedRecipe(null)}
        />
      )}

      <RecipeForm
        open={showForm}
        initialData={editingRecipe ?? undefined}
        onCancel={closeForm}
        onSubmit={(data) => handleSubmit(data)}
      />

      {deletingRecipe && (
        <ConfirmDialog
          open={!!deletingRecipe}
          title="Delete this recipe?"
          message={`Are you sure you want to delete "${deletingRecipe.name}"?`}
          onCancel={() => setDeletingRecipe(null)}
          onConfirm={async () => {
            await fetch(
              `${appConfig.baseApiUrl}/recipes/${deletingRecipe.id}`,
              {
                method: "DELETE",
                headers: { Authorization: `Bearer ${token}` },
              }
            );
            setDeletingRecipe(null);
            fetchRecipes(page, limit, search);
          }}
        />
      )}
    </div>
  );
};

export default Recipes;
