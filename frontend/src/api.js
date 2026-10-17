import axios from 'axios';

const api = axios.create({ baseURL: import.meta.env.VITE_API_URL || '/api' });

export const getCategories = () => api.get('/categories');
export const getMenu = (params) => api.get('/menu', { params });
export const getDish = (id) => api.get(`/menu/${id}`);
export const createDish = (data) => api.post('/menu', data);
export const updateDish = (id, data) => api.put(`/menu/${id}`, data);
export const deleteDish = (id) => api.delete(`/menu/${id}`);

export const uploadImage = (file, alt = '') => {
  const form = new FormData();
  form.append('file', file);
  form.append('alt', alt);
  return api.post('/images', form);
};

export const errorText = (e) =>
  e?.response?.data?.error || e?.message || 'Неизвестная ошибка';

export default api;
