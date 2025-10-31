export class BasePagination {
  public limit: number;
  public totalDocs: number;
  public nextPage: number | null;
  public hasNextPage: boolean;
  public hasPrevPage: boolean;
  public prevPage: number | null;
  public page: number;
  public totalPages?: number;

  public constructor({
    limit,
    page,
    totalDocs,
  }: Pick<BasePagination, "limit" | "page" | "totalDocs">) {
    this.totalDocs = totalDocs;
    this.limit = limit;
    this.page = page;

    const totalPages = Math.ceil(this.totalDocs / this.limit);
    this.hasNextPage = this.page >= 1 && this.page < totalPages;
    this.hasPrevPage = this.page > 1 && this.page <= totalPages + 1;
    this.prevPage = this.hasPrevPage ? this.page - 1 : null;
    this.nextPage = this.hasNextPage ? this.page + 1 : null;
    this.totalPages = totalPages;
  }
}
