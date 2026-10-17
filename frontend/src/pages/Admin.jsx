import { Link } from 'react-router-dom';
import { deleteDish, errorText } from '../api';
import DishCard from '../components/DishCard';
import useMenuData from '../hooks/useMenuData';

// Экран 2: редактирование меню - все блюда по категориям, изменение/удаление и добавление в категорию.
export default function Admin() {
  const { cats, items, loading, error, setError, reload } = useMenuData();

  const remove = async (dish) => {
    if (!window.confirm(`Удалить «${dish.name}»?`)) return;
    try {
      await deleteDish(dish.id);
      reload();
    } catch (e) {
      setError(errorText(e));
    }
  };

  return (
    <>
      <h1 className="title">Редактирование меню</h1>
      {error && <div className="alert">{error}</div>}
      {loading && <p>Загрузка…</p>}
      {cats.map((cat) => {
        const list = items.filter((i) => i.category_id === cat.id);
        return (
          <section key={cat.id} className="section">
            <div className="section-head">
              <h3>{cat.name}</h3>
              <Link className="btn" to={`/dish/new?category=${cat.id}`}>+ Добавить позицию</Link>
            </div>
            {list.length === 0 ? (
              <p className="muted">В этой категории пока нет позиций</p>
            ) : (
              <div className="grid">
                {list.map((d) => <DishCard key={d.id} dish={d} editable onDelete={remove} />)}
              </div>
            )}
          </section>
        );
      })}
    </>
  );
}
