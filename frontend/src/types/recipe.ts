import type { BasePagination } from "../lib/pagination";

export type RecipeDto = {
  id: string;
  name: string;
  description?: string;
  picture?: string;
  ingredients: Record<string, string>;
  createdAt: string;
  updatedAt: string;
};

export type GetRecipesResponseDto = {
  recipes: RecipeDto[];
  meta: BasePagination;
};
