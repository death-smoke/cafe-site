import { useEffect, useState } from 'react';
import { errorText, getCategories, getMenu } from '../api';
import DishCard from '../components/DishCard';

// Экран 3: поиск, фильтры по категориям, состав и описание (без изменения и удаления).
export default function Search() {
  const [cats, setCats] = useState([]);
  const [category, setCategory] = useState('');
  const [query, setQuery] = useState('');
  const [items, setItems] = useState([]);
  const [selected, setSelected] = useState(null);
  const [error, setError] = useState('');

  useEffect(() => {
    getCategories().then((r) => setCats(r.data)).catch((e) => setError(errorText(e)));
  }, []);

  useEffect(() => {
    const t = setTimeout(() => {
      const params = { limit: 500, sort: 'name' };
      if (category) params.category_id = category;
      if (query.trim()) params.search = query.trim();
      getMenu(params)
        .then((r) => { setItems(r.data.items); setError(''); })
        .catch((e) => setError(errorText(e)));
    }, 250);
    return () => clearTimeout(t);
  }, [query, category]);

  return (
    <>
      <h2>Поиск по меню</h2>
      <input
        className="search"
        type="search"
        placeholder="Название, описание или ингредиент…"
        value={query}
        onChange={(e) => setQuery(e.target.value)}
      />
      <div className="chips">
        <button className={`chip ${category === '' ? 'active' : ''}`} onClick={() => setCategory('')}>Все</button>
        {cats.map((c) => (
          <button key={c.id} className={`chip ${category === c.id ? 'active' : ''}`} onClick={() => setCategory(c.id)}>
            {c.name}
          </button>
        ))}
      </div>
      {error && <div className="alert">{error}</div>}
      <p className="muted">Найдено: {items.length}</p>
      <div className="grid">
        {items.map((d) => <DishCard key={d.id} dish={d} onDetails={setSelected} />)}
      </div>

      {selected && (
        <div className="modal-bg" onClick={() => setSelected(null)}>
          <div className="modal" onClick={(e) => e.stopPropagation()}>
            <button className="close" onClick={() => setSelected(null)} aria-label="Закрыть">×</button>
            {selected.image_url && <img src={selected.image_url} alt={selected.name} />}
            <h3>{selected.name} — {selected.price} ₽</h3>
            <p className="muted">{selected.category_name}</p>
            <h4>Описание</h4>
            <p>{selected.description || '—'}</p>
            <h4>Состав</h4>
            <p>{selected.ingredients || '—'}</p>
          </div>
        </div>
      )}
    </>
  );
}
