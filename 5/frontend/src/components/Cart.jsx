export default function Cart({ items }) {
  const total = items.reduce((sum, item) => sum + item.price, 0);

  return (
    <section>
      <h2>Koszyk</h2>
      {items.length === 0 ? (
        <p>Koszyk jest pusty</p>
      ) : (
        <ul>
          {items.map((item, index) => (
            <li key={`${item.id}-${index}`}>
              {item.name} - {item.price.toFixed(2)} PLN
            </li>
          ))}
        </ul>
      )}
      <p>Razem: {total.toFixed(2)} PLN</p>
    </section>
  );
}
