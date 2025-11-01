export type BasePagination = {
  limit: number;
  totalItems: number;
  nextPage: number | null;
  hasNextPage: boolean;
  hasPrevPage: boolean;
  prevPage: number | null;
  page: number;
  totalPages: number;
};
