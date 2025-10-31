import { useEffect, useState, useCallback } from "react";
import { RecipeCard } from "@/components/recipes/RecipeCard";
import { RecipeDrawer } from "@/components/recipes/RecipeDrawer";
import { Loader2, Search } from "lucide-react";
import { Button } from "@/components/ui/button";
import debounce from "lodash.debounce";
import {
  InputGroup,
  InputGroupAddon,
  InputGroupInput,
} from "../components/ui/input-group";
import type { GetRecipesResponseDto, RecipeDto } from "../types/recipe";
import { appConfig } from "../lib/config";
import { useAuth } from "../context/AuthContext";
import { Topbar } from "../components/layout/Topbar";

const LIMIT_OPTIONS = [5, 10, 20] as const;

const Recipes = () => {
  const { token, logout, user } = useAuth();

  const [recipes, setRecipes] = useState<RecipeDto[]>([]);
  const [loading, setLoading] = useState(true);

  const [page, setPage] = useState(1);
  const [limit, setLimit] = useState(10);
  const [totalPages, setTotalPages] = useState(1);

  const [search, setSearch] = useState("");
  const [selectedRecipe, setSelectedRecipe] = useState<RecipeDto | null>(null);

  const debouncedFetch = useCallback(
    debounce((page, limit, search) => {
      fetchRecipes(page, limit, search);
    }, 200),
    []
  );

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
      setTotalPages(data?.meta?.totalPages || 2); // todo: add meta to backend
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    debouncedFetch(page, limit, search);
  }, [page, limit, search]);

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
            <RecipeCard
              key={recipe.id}
              recipe={recipe}
              onDetails={() => setSelectedRecipe(recipe)}
            />
          ))}
        </div>
      )}

      {selectedRecipe && (
        <RecipeDrawer
          recipe={selectedRecipe}
          onClose={() => setSelectedRecipe(null)}
        />
      )}
    </div>
  );
};

export default Recipes;
