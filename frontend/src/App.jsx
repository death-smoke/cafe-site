import { NavLink, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import Admin from './pages/Admin';
import DishForm from './pages/DishForm';
import Search from './pages/Search';

export default function App() {
  return (
    <>
      <header className="navbar">
        <span className="brand"><img src="/logo.png" alt=""/>Солнышко</span>
        <nav>
          <NavLink to="/" end>Меню</NavLink>
          <NavLink to="/search">Поиск</NavLink>
          <NavLink to="/admin">Редактирование</NavLink>
        </nav>
      </header>
      <main className="container">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/search" element={<Search />} />
          <Route path="/admin" element={<Admin />} />
          <Route path="/dish/new" element={<DishForm />} />
          <Route path="/dish/:id/edit" element={<DishForm />} />
          <Route path="*" element={<p>Страница не найдена</p>} />
        </Routes>
      </main>
    </>
  );
}
