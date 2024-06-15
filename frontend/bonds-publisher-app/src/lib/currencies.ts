export enum Currency {
  USD = "USD",
  MXN = "MXN",
}

export interface Location {
  label: string;
  locale: string;
}

export const locations: Record<Currency, Location> = {
  [Currency.USD]: { label: "$ Dollar", locale: "en-US" },
  [Currency.MXN]: { label: "$ Peso Mexicano", locale: "es-MX" },
};
