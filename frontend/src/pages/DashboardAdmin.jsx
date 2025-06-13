import React, { useState, useEffect } from "react";
import "./DashboardAdmin.css";
import Swal from "sweetalert2";
import { useNavigate } from "react-router-dom";

function DashboardAdmin() {
  const [nombre, setNombre] = useState("");
  const [descripcion, setDescripcion] = useState("");
  const [profesor, setProfesor] = useState("");
  const [horario, setHorario] = useState("");
  const [cupo, setCupo] = useState("");
  const [actividades, setActividades] = useState([]);
  const [modoEdicion, setModoEdicion] = useState(false);
  const [idEditar, setIdEditar] = useState(null);

  const token = localStorage.getItem("token");
  const navigate = useNavigate();

  const cargarActividades = () => {
    fetch("http://localhost:8080/actividades")
      .then(res => res.json())
      .then(data => setActividades(data))
      .catch(err => console.error("Error al cargar actividades", err));
  };

  useEffect(() => {
    cargarActividades();
  }, []);

  const limpiarFormulario = () => {
    setNombre("");
    setDescripcion("");
    setProfesor("");
    setHorario("");
    setCupo("");
    setModoEdicion(false);
    setIdEditar(null);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const actividad = {
      nombre,
      descripcion,
      profesor,
      horario,
      cupo: parseInt(cupo),
      categoria: "General",
      disponible: parseInt(cupo),
      estado: true
    };

    try {
      const url = idEditar
        ? `http://localhost:8080/actividades/${idEditar}`
        : "http://localhost:8080/actividades";
      const metodo = idEditar ? "PUT" : "POST";

      const res = await fetch(url, {
        method: metodo,
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`
        },
        body: JSON.stringify(actividad)
      });

      const data = await res.json();

      if (!res.ok) throw new Error(data.error || "Error desconocido");

      Swal.fire(
        idEditar ? "Actividad actualizada" : "Actividad creada",
        "Operación exitosa",
        "success"
      );
      cargarActividades();
      limpiarFormulario();
    } catch (err) {
      Swal.fire("Error", err.message, "error");
    }
  };

  const handleEditar = (act) => {
    setModoEdicion(true);
    setIdEditar(act.id);
    setNombre(act.nombre);
    setDescripcion(act.descripcion);
    setProfesor(act.profesor);
    setHorario(act.horario);
    setCupo(act.cupo);
  };

  const handleEliminar = (id) => {
    Swal.fire({
      title: "¿Estás seguro?",
      text: "Esta acción eliminará la actividad.",
      icon: "warning",
      showCancelButton: true,
      confirmButtonText: "Sí, eliminar",
      cancelButtonText: "Cancelar"
    }).then(async (result) => {
      if (result.isConfirmed) {
        try {
          const res = await fetch(`http://localhost:8080/actividades/${id}`, {
            method: "DELETE",
            headers: { Authorization: `Bearer ${token}` }
          });

          if (!res.ok) {
            const data = await res.json();
            throw new Error(data.error || "Error al eliminar");
          }

          Swal.fire("Eliminada", "Actividad eliminada correctamente", "success");
          cargarActividades();
        } catch (err) {
          Swal.fire("Error", err.message, "error");
        }
      }
    });
  };

  return (
    <main className="admin-container">
      {/* Botón volver */}
      <button className="volver-btn" onClick={() => navigate(-1)}>
        ← Volver
      </button>

      <h1 className="admin-titulo">
        {modoEdicion ? "Editar Actividad" : "Agregar Actividad"}
      </h1>

      <form className="admin-formulario" onSubmit={handleSubmit}>
        <input type="text" placeholder="Nombre" value={nombre} onChange={(e) => setNombre(e.target.value)} required />
        <textarea placeholder="Descripción" value={descripcion} onChange={(e) => setDescripcion(e.target.value)} required />
        <input type="text" placeholder="Profesor" value={profesor} onChange={(e) => setProfesor(e.target.value)} required />
        <input type="text" placeholder="Horario" value={horario} onChange={(e) => setHorario(e.target.value)} required />
        <input type="number" placeholder="Cupo" value={cupo} onChange={(e) => setCupo(e.target.value)} required />
        <button type="submit">
          {modoEdicion ? "Guardar Cambios" : "Agregar Actividad"}
        </button>
        {modoEdicion && (
          <button type="button" className="admin-btn eliminar" onClick={limpiarFormulario}>
            Cancelar edición
          </button>
        )}
      </form>

      <section className="admin-actividades">
        <h2>Actividades Cargadas</h2>
        <ul className="admin-lista">
          {actividades.length === 0 && <p>No hay actividades registradas aún.</p>}
          {actividades.map((a) => (
            <li key={a.id}>
              <strong>{a.nombre}</strong> - {a.horario} - Prof: {a.profesor}
              <div className="admin-botones">
                <button className="admin-btn editar" onClick={() => handleEditar(a)}>Editar</button>
                <button className="admin-btn eliminar" onClick={() => handleEliminar(a.id)}>Eliminar</button>
              </div>
            </li>
          ))}
        </ul>
      </section>
    </main>
  );
}

export default DashboardAdmin;
