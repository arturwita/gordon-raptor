export const UserRole = {
  User: "USER",
  Admin: "ADMIN",
} as const;

export type UserRole = (typeof UserRole)[keyof typeof UserRole];

export type UserJwtPayload = {
  sub: string;
  email: string;
  firstName?: string;
  lastName?: string;
  picture?: string;
  role: UserRole;
  exp: number;
  iat: number;
};
