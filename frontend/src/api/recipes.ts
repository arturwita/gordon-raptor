import type { RecipeDto, GetRecipesResponseDto } from "../types/recipe";

export async function fetchRecipes(): Promise<RecipeDto[]> {
  const res = await fetch("http://localhost:8000/recipes");

  if (!res.ok) {
    throw new Error(`Failed to fetch recipes: ${res.status}`);
  }

  const response: GetRecipesResponseDto = await res.json();

  return response.recipes;
}
