function isRecord(value) {
  return typeof value === "object" && value !== null && !Array.isArray(value);
}

export function parseProducts(data) {
  if (!Array.isArray(data)) {
    return null;
  }

  const products = data.filter(
    (item) =>
      isRecord(item) &&
      Number.isInteger(item.id) &&
      typeof item.name === "string" &&
      typeof item.price === "number" &&
      Number.isFinite(item.price),
  );

  return products.length === data.length ? products : null;
}

export function parseMessage(data) {
  if (!isRecord(data) || typeof data.message !== "string") {
    return null;
  }
  return data.message;
}
