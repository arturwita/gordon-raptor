import type { BasePagination } from "./pagination";

export type RecipeDto = {
  id: string;
  name: string;
  description?: string;
  picture?: string;
  ingredients: Record<string, string>;
  createdAt: string;
  updatedAt: string;
};

export type RecipeInput = Omit<RecipeDto, "id" | "createdAt" | "updatedAt">;

export type GetRecipesResponseDto = {
  recipes: RecipeDto[];
  meta: BasePagination;
};
