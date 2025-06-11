// DashboardAdmin.jsx
import React, { useState, useEffect } from "react";
import "./DashboardAdmin.css";
import Swal from "sweetalert2";

function DashboardAdmin() {
  const [titulo, setTitulo] = useState("");
  const [descripcion, setDescripcion] = useState("");
  const [profesor, setProfesor] = useState("");
  const [horario, setHorario] = useState("");
  const [duracion, setDuracion] = useState(""); // Este campo no se usa en el backend
  const [cupo, setCupo] = useState("");
  const [actividades, setActividades] = useState([]);

  const token = localStorage.getItem("token");

  useEffect(() => {
    fetch("http://localhost:8080/actividades")
      .then(res => res.json())
      .then(data => setActividades(data))
      .catch(err => console.error("Error al cargar actividades", err));
  }, []);

  const handleSubmit = (e) => {
    e.preventDefault();

    const nuevaActividad = {
      nombre: titulo,
      descripcion,
      profesor,
      horario,
      cupo: parseInt(cupo),
      categoria: "General",
      disponible: parseInt(cupo),
      estado: true
    };

    fetch("http://localhost:8080/actividades", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`
      },
      body: JSON.stringify(nuevaActividad)
    })
      .then(res => {
        if (!res.ok) throw new Error("Error al crear actividad");
        return res.json();
      })
      .then(data => {
        Swal.fire("Actividad creada", "Se agregó correctamente", "success");
        setActividades([...actividades, data]);
        setTitulo("");
        setDescripcion("");
        setProfesor("");
        setHorario("");
        setDuracion("");
        setCupo("");
      })
      .catch(err => {
        Swal.fire("Error", "No se pudo crear la actividad", "error");
        console.error(err);
      });
  };

  return (
    <main className="admin-container">
      <h1 className="admin-titulo">Panel de Administración</h1>
      <form className="admin-form" onSubmit={handleSubmit}>
        <input type="text" placeholder="Título" value={titulo} onChange={(e) => setTitulo(e.target.value)} required />
        <textarea placeholder="Descripción" value={descripcion} onChange={(e) => setDescripcion(e.target.value)} required />
        <input type="text" placeholder="Profesor" value={profesor} onChange={(e) => setProfesor(e.target.value)} required />
        <input type="text" placeholder="Horario" value={horario} onChange={(e) => setHorario(e.target.value)} required />
        <input type="text" placeholder="Duración (opcional)" value={duracion} onChange={(e) => setDuracion(e.target.value)} />
        <input type="number" placeholder="Cupo" value={cupo} onChange={(e) => setCupo(e.target.value)} required />
        <button type="submit">Agregar Actividad</button>
      </form>

      <section className="admin-actividades">
        <h2>Actividades Cargadas</h2>
        <ul>
          {actividades.length === 0 && <p>No hay actividades registradas aún.</p>}
          {actividades.map((a) => (
            <li key={a.id}>
              <strong>{a.nombre}</strong> - {a.horario} - Prof: {a.profesor}
            </li>
          ))}
        </ul>
      </section>
    </main>
  );
}

export default DashboardAdmin;
