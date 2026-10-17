import { useEffect, useState } from 'react';
import { useNavigate, useParams, useSearchParams } from 'react-router-dom';
import { createDish, deleteDish, errorText, getCategories, getDish, updateDish, uploadImage } from '../api';

const emptyForm = (category = '') => ({ name: '', category_id: category, price: '', description: '', ingredients: '' });

export default function DishForm() {
  const { id } = useParams();
  const [params] = useSearchParams();
  const presetCategory = params.get('category') || '';
  const navigate = useNavigate();
  const back = () => navigate('/admin');

  const [form, setForm] = useState(emptyForm(presetCategory));
  const [cats, setCats] = useState([]);
  const [imageId, setImageId] = useState(null);
  const [preview, setPreview] = useState('');
  const [file, setFile] = useState(null);
  const [error, setError] = useState('');
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    getCategories().then((r) => setCats(r.data)).catch((e) => setError(errorText(e)));
  }, []);

  useEffect(() => {
    if (!id) {
      setForm(emptyForm(presetCategory)); setImageId(null); setPreview(''); setFile(null);
      return;
    }
    getDish(id)
      .then((r) => {
        const d = r.data;
        setForm({
          name: d.name, category_id: d.category_id, price: d.price,
          description: d.description, ingredients: d.ingredients,
        });
        setImageId(d.image_id);
        setPreview(d.image_url || '');
      })
      .catch((e) => setError(errorText(e)));
  }, [id, presetCategory]);

  const set = (k) => (e) => setForm({ ...form, [k]: e.target.value });

  const onFile = (e) => {
    const f = e.target.files[0];
    if (!f) return;
    if (f.type !== 'image/png') {
      setError('Допустим только формат PNG');
      e.target.value = '';
      return;
    }
    setError('');
    setFile(f);
    setPreview(URL.createObjectURL(f));
  };

  const removeImage = () => { setFile(null); setImageId(null); setPreview(''); };

  const submit = async (e) => {
    e.preventDefault();
    setError('');
    if (!form.name.trim()) return setError('Укажите название');
    if (!form.category_id) return setError('Выберите категорию');
    if (!(Number(form.price) > 0)) return setError('Цена должна быть больше нуля');
    setSaving(true);
    try {
      let imgId = imageId;
      if (file) {
        const up = await uploadImage(file, form.name);
        imgId = up.data.id;
      }
      const payload = {
        name: form.name.trim(),
        category_id: Number(form.category_id),
        price: Number(form.price),
        description: form.description,
        ingredients: form.ingredients,
        image_id: imgId,
      };
      if (id) await updateDish(id, payload);
      else await createDish(payload);
      back();
    } catch (err) {
      setError(errorText(err));
    } finally {
      setSaving(false);
    }
  };

  const remove = async () => {
    if (!window.confirm('Удалить это блюдо?')) return;
    try { await deleteDish(id); back(); } catch (err) { setError(errorText(err)); }
  };

  return (
    <form className="form" onSubmit={submit}>
      <h2>{id ? 'Редактирование блюда' : 'Новая позиция меню'}</h2>
      {error && <div className="alert">{error}</div>}

      <label>Название
        <input value={form.name} onChange={set('name')} maxLength={120} required />
      </label>
      <label>Категория
        <select value={form.category_id} onChange={set('category_id')} required>
          <option value="">— выберите —</option>
          {cats.map((c) => <option key={c.id} value={c.id}>{c.name}</option>)}
        </select>
      </label>
      <label>Цена, ₽
        <input type="number" min="1" step="1" value={form.price} onChange={set('price')} required />
      </label>
      <label>Описание
        <textarea rows={3} value={form.description} onChange={set('description')} maxLength={2000} />
      </label>
      <label>Состав
        <textarea rows={3} value={form.ingredients} onChange={set('ingredients')} maxLength={2000}
          placeholder="Через запятую: курица, сыр, соус…" />
      </label>
      <label>Изображение (PNG)
        <input type="file" accept="image/png" onChange={onFile} />
      </label>
      {preview && (
        <div className="preview">
          <img src={preview} alt="Предпросмотр" />
          <button type="button" className="btn btn-outline" onClick={removeImage}>Убрать картинку</button>
        </div>
      )}

      <div className="actions">
        <button className="btn" disabled={saving}>{saving ? 'Сохранение…' : 'Сохранить'}</button>
        <button type="button" className="btn btn-outline" onClick={back}>Отмена</button>
        {id && <button type="button" className="btn btn-danger" onClick={remove}>Удалить</button>}
      </div>
    </form>
  );
}
