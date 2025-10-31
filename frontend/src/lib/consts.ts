export const Theme = {
  Dark: "dark",
  Light: "light",
} as const;

export type ThemeType = (typeof Theme)[keyof typeof Theme];
