import React from "react";
import { useParams, useNavigate } from "react-router-dom";
import './ActividadDetalle.css';

// Simulación de datos
const actividadesFake = [
  {
    id: 1,
    titulo: "Fútbol",
    descripcion: "Partido amistoso de fútbol 5",
    profesor: "Carlos Pérez",
    horario: "18:00 a 20:00",
    duracion: "2 horas",
    cupo: 10,
    imagen: "/imagenes/futbol.jpg"
  },
  {
    id: 2,
    titulo: "Yoga",
    descripcion: "Clase de yoga para principiantes",
    profesor: "Laura Gómez",
    horario: "09:00 a 11:00",
    duracion: "2 horas",
    cupo: 15,
    imagen: "/imagenes/yoga.jpg"
  },
  {
    id: 3,
    titulo: "Crossfit",
    descripcion: "Entrenamiento de alta intensidad",
    profesor: "Martín Díaz",
    horario: "20:00 a 22:00",
    duracion: "2 horas",
    cupo: 12,
    imagen: "/imagenes/crossfit.jpg"
  }
];

function ActividadDetalle() {
  const { id } = useParams();
  const navigate = useNavigate();
  const actividad = actividadesFake.find((a) => a.id === parseInt(id));

  if (!actividad) return <div style={{ padding: "2rem", color: "white" }}>Actividad no encontrada</div>;

  return (
    <main className="detalle-container">
      <button onClick={() => navigate(-1)} className="volver-btn">← Volver</button>
      <h1 className="detalle-titulo">{actividad.titulo}</h1>
      <img src={actividad.imagen} alt={actividad.titulo} className="detalle-imagen" />
      <p><strong>Profesor:</strong> {actividad.profesor}</p>
      <p><strong>Horario:</strong> {actividad.horario}</p>
      <p><strong>Duración:</strong> {actividad.duracion}</p>
      <p><strong>Cupo:</strong> {actividad.cupo}</p>
      <p><strong>Descripción:</strong> {actividad.descripcion}</p>
    </main>
  );
}

export default ActividadDetalle;
