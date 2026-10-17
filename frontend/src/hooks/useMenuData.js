import { useCallback, useEffect, useState } from 'react';
import { errorText, getCategories, getMenu } from '../api';

// Загружает категории и все блюда; reload() перечитывает данные.
export default function useMenuData() {
  const [cats, setCats] = useState([]);
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  const reload = useCallback(async () => {
    try {
      const [c, m] = await Promise.all([getCategories(), getMenu({ limit: 500, sort: 'name' })]);
      setCats(c.data);
      setItems(m.data.items);
      setError('');
    } catch (e) {
      setError(errorText(e));
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => { reload(); }, [reload]);

  return { cats, items, loading, error, setError, reload };
}
