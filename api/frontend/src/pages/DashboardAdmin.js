import React, { useState } from "react";
import './DashboardAdmin.css';

const AdminActividades = () => {
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
    <main className="container">
      <h1 className="titulo">Panel de Administración</h1>

      <section className="formSection">
        <h2>Agregar / Editar Actividad</h2>
        <form onSubmit={handleSubmit} className="form">
          <label>Nombre de la actividad</label>
          <input
            type="text"
            value={nombre}
            onChange={(e) => setNombre(e.target.value)}
            placeholder="Nombre actividad"
            required
            className="input"
          />

          <label>Descripción</label>
          <textarea
            rows="4"
            value={descripcion}
            onChange={(e) => setDescripcion(e.target.value)}
            placeholder="Descripción de la actividad"
            required
            className="textarea"
          />

          <button type="submit" className="btn">
            {editIndex !== null ? "Guardar Cambios" : "Agregar Actividad"}
          </button>
        </form>
      </section>

      <section>
        <h2>Actividades existentes</h2>
        <ul className="lista">
          {actividades.length === 0 && <li>No hay actividades aún</li>}
          {actividades.map((act, i) => (
            <li key={i} className="itemLista">
              <div>
                <strong>{act.nombre}</strong>: {act.descripcion}
              </div>
              <div>
                <button onClick={() => handleEditar(i)} className="btnEdit">
                  Editar
                </button>
                <button onClick={() => handleEliminar(i)} className="btnDelete">
                  Eliminar
                </button>
              </div>
            </li>
          ))}
        </ul>
      </section>
    </main>
  );
};

export default AdminActividades;
