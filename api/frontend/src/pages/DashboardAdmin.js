import React, { useState } from "react";
import './DashboardAdmin.css';

const DashboardAdmin = () => {
  const [actividades, setActividades] = useState([]);
  const [nombre, setNombre] = useState("");
  const [descripcion, setDescripcion] = useState("");
  const [editIndex, setEditIndex] = useState(null);

  const handleSubmit = (e) => {
    e.preventDefault();
    if (!nombre.trim() || !descripcion.trim()) return;

    if (editIndex !== null) {
      const nuevas = [...actividades];
      nuevas[editIndex] = { nombre, descripcion };
      setActividades(nuevas);
      setEditIndex(null);
    } else {
      setActividades([...actividades, { nombre, descripcion }]);
    }

    setNombre("");
    setDescripcion("");
  };

  const handleEditar = (index) => {
    setNombre(actividades[index].nombre);
    setDescripcion(actividades[index].descripcion);
    setEditIndex(index);
  };

  const handleEliminar = (index) => {
    setActividades(actividades.filter((_, i) => i !== index));
    if (editIndex === index) {
      setNombre("");
      setDescripcion("");
      setEditIndex(null);
    }
  };

  return (
    <main className="admin-container">
      <h1 className="admin-titulo">Panel de Administración</h1>

      <section className="admin-formSection">
        <h2 className="admin-subtitulo">Agregar / Editar Actividad</h2>
        <form onSubmit={handleSubmit} className="admin-form">
          <label>Nombre de la actividad: </label>
          <input
            type="text"
            value={nombre}
            onChange={(e) => setNombre(e.target.value)}
            placeholder="Nombre actividad"
            required
            className="admin-input"
          />

          <label>Descripción: </label>
          <textarea
            rows="4"
            value={descripcion}
            onChange={(e) => setDescripcion(e.target.value)}
            placeholder="Descripción de la actividad"
            required
            className="admin-textarea"
          />

          <button type="submit" className="admin-btn">
            {editIndex !== null ? "Guardar Cambios" : "Agregar Actividad"}
          </button>
        </form>
      </section>

      <section>
        <h2 className="admin-subtitulo">Actividades existentes</h2>
        <ul className="admin-lista">
          {actividades.length === 0 && <li>No hay actividades aún</li>}
          {actividades.map((act, i) => (
            <li key={i} className="admin-itemLista">
              <div>
                <strong>{act.nombre}</strong>: {act.descripcion}
              </div>
              <div>
                <button onClick={() => handleEditar(i)} className="admin-btnEdit">Editar</button>
                <button onClick={() => handleEliminar(i)} className="admin-btnDelete">Eliminar</button>
              </div>
            </li>
          ))}
        </ul>
      </section>
    </main>
  );
};

export default DashboardAdmin;

