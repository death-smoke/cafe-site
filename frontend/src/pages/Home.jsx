import DishCard from '../components/DishCard';
import useMenuData from '../hooks/useMenuData';

// Экран 1: просмотр меню (без редактирования и удаления).
export default function Home() {
  const { cats, items, loading, error } = useMenuData();

  return (
    <>
      <h1 className="title">{"Кафе 'Солнышко'"}</h1>
      <h2 className="subtitle">Меню</h2>
      {error && <div className="alert">{error}</div>}
      {loading && <p>Загрузка…</p>}
      {cats.map((cat) => {
        const list = items.filter((i) => i.category_id === cat.id);
        return (
          <section key={cat.id} className="section">
            <h3>{cat.name}</h3>
            {list.length === 0 ? (
              <p className="muted">В этой категории пока нет позиций</p>
            ) : (
              <div className="grid">
                {list.map((d) => <DishCard key={d.id} dish={d} />)}
              </div>
            )}
          </section>
        );
      })}
    </>
  );
}
