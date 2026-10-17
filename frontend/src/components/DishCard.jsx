import { Link } from 'react-router-dom';

export default function DishCard({ dish, onDetails, editable = false, onDelete }) {
  const attrs = dish.image_attributes || {};
  const hasActions = onDetails || editable || onDelete;
  return (
    <article className="card">
      <div className="card-img">
        {dish.image_url ? (
          <img
            src={dish.image_url}
            alt={attrs.alt || dish.name}
            width={attrs.width}
            height={attrs.height}
            loading="lazy"
          />
        ) : (
          <span className="placeholder"><img src="/placeholder.png" alt=""/></span>
        )}
      </div>
      <div className="card-body">
        <div className="card-title">
          <h4>{dish.name}</h4>
          <span className="price">{dish.price} ₽</span>
        </div>
        <p className="muted">{dish.description || 'Без описания'}</p>
        {hasActions && (
          <div className="actions">
            {onDetails && (
              <button className="btn" onClick={() => onDetails(dish)}>Состав и описание</button>
            )}
            {editable && (
              <Link className="btn btn-outline" to={`/dish/${dish.id}/edit`}>Изменить</Link>
            )}
            {onDelete && (
              <button className="btn btn-danger" onClick={() => onDelete(dish)}>Удалить</button>
            )}
          </div>
        )}
      </div>
    </article>
  );
}
