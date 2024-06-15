import { Currency, locations } from "./currencies";

export function formatterForCurrency(currency: Currency) {
  const locale = locations[currency].locale;

  return new Intl.NumberFormat(locale, {
    style: "currency",
    currency,
    minimumFractionDigits: 4,
    maximumFractionDigits: 4,
  });
}
