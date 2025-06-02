import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import Swal from "sweetalert2";
import './DashboardAdmin.css';

function DashboardAdmin() {
  const navigate = useNavigate();

  const [titulo, setTitulo] = useState("");
  const [horario, setHorario] = useState("");
  const [profesor, setProfesor] = useState("");
  const [descripcion, setDescripcion] = useState("");

  const [actividades, setActividades] = useState([]);
  const [modoEdicion, setModoEdicion] = useState(false);
  const [actividadEditando, setActividadEditando] = useState(null);

  const handleAgregarActividad = () => {
    if (!titulo || !horario || !profesor || !descripcion) {
      Swal.fire({
        icon: 'warning',
        title: 'Campos incompletos',
        text: 'Por favor completá todos los campos',
      });
      return;
    }

    const nuevaActividad = {
      id: actividades.length + 1,
      titulo,
      horario,
      profesor,
      descripcion,
    };

    setActividades([...actividades, nuevaActividad]);

    Swal.fire({
      icon: 'success',
      title: 'Actividad agregada',
      text: `Se agregó "${titulo}" correctamente.`,
    });

    limpiarFormulario();
  };

  const limpiarFormulario = () => {
    setTitulo("");
    setHorario("");
    setProfesor("");
    setDescripcion("");
    setModoEdicion(false);
    setActividadEditando(null);
  };

  const handleEliminar = (id) => {
    const filtradas = actividades.filter((act) => act.id !== id);
    setActividades(filtradas);
    Swal.fire("Eliminado", "La actividad fue eliminada", "success");
  };

  const handleEditar = (actividad) => {
    setTitulo(actividad.titulo);
    setHorario(actividad.horario);
    setProfesor(actividad.profesor);
    setDescripcion(actividad.descripcion);
    setModoEdicion(true);
    setActividadEditando(actividad);
  };

  const handleActualizarActividad = () => {
    const actualizadas = actividades.map((act) =>
      act.id === actividadEditando.id
        ? { ...act, titulo, horario, profesor, descripcion }
        : act
    );
    setActividades(actualizadas);
    Swal.fire("Actualizado", "La actividad fue editada correctamente", "success");
    limpiarFormulario();
  };

  return (
    <main className="admin-container">
      <button onClick={() => navigate(-1)} className="volver-btn">← Volver</button>
      <h1 className="admin-titulo">Panel del Administrador</h1>

      <div className="admin-grid">
        {/* Columna izquierda: Agregar */}
        <section className="admin-card">
          <h2>Agregar Actividades</h2>
          <form className="admin-formulario" onSubmit={(e) => e.preventDefault()}>
            <input
              type="text"
              placeholder="Título de la actividad"
              value={titulo}
              onChange={(e) => setTitulo(e.target.value)}
            />
            <input
              type="text"
              placeholder="Horario"
              value={horario}
              onChange={(e) => setHorario(e.target.value)}
            />
            <input
              type="text"
              placeholder="Profesor"
              value={profesor}
              onChange={(e) => setProfesor(e.target.value)}
            />
            <textarea
              placeholder="Descripción"
              value={descripcion}
              onChange={(e) => setDescripcion(e.target.value)}
            ></textarea>
            <button onClick={modoEdicion ? handleActualizarActividad : handleAgregarActividad}>
              {modoEdicion ? "Actualizar Actividad" : "Agregar Actividad"}
            </button>
          </form>
        </section>

        {/* Columna derecha: Editar */}
        <section className="admin-card">
          <h2>Editar Actividad</h2>
          <h3>Actividades Disponibles</h3>
          <ul className="admin-lista">
            {actividades.map((act) => (
              <li key={act.id}>
                <strong>{act.titulo}</strong> - {act.horario} - {act.profesor}<br />
                <span>{act.descripcion}</span>
                <div className="admin-botones">
                  <button onClick={() => handleEditar(act)} className="admin-btn editar">Editar</button>
                  <button onClick={() => handleEliminar(act.id)} className="admin-btn eliminar">Eliminar</button>
                </div>
              </li>
            ))}
          </ul>
        </section>
      </div>
    </main>
  );
}

export default DashboardAdmin;
