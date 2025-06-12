import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import Swal from "sweetalert2";
import "./ActividadDetalle.css";

function ActividadDetalle() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [actividad, setActividad] = useState(null);
  const token = localStorage.getItem("token");

  useEffect(() => {
    fetch(`http://localhost:8080/api/actividades/${id}`, {
      headers: { Authorization: `Bearer ${token}` }
    })
      .then((res) => res.json())
      .then((data) => setActividad(data))
      .catch((err) => {
        console.error("Error al obtener actividad:", err);
        Swal.fire("Error", "No se pudo cargar la actividad", "error");
      });
  }, [id, token]);

  const handleInscribirse = () => {
    fetch("http://localhost:8080/inscripciones", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`
      },
      body: JSON.stringify({ actividad_id: actividad.id })
    })
      .then(res => res.json())
      .then(data => {
        Swal.fire("¡Inscripción exitosa!", data.mensaje, "success");
      })
      .catch(err => {
        Swal.fire("Error", "No se pudo realizar la inscripción", "error");
        console.error(err);
      });
  };

  if (!actividad) {
    return <p>Cargando actividad...</p>;
  }

  return (
    <div className="actividad-detalle">
      <button onClick={() => navigate(-1)} className="volver-btn">← Volver</button>
      <h2>{actividad.titulo}</h2>
      <img src={actividad.foto || "/default.jpg"} alt={actividad.titulo} className="actividad-img" />
      <p><strong>Descripción:</strong> {actividad.descripcion}</p>
      <p><strong>Profesor:</strong> {actividad.profesor}</p>
      <p><strong>Horario:</strong> {actividad.horario}</p>
      <p><strong>Duración:</strong> {actividad.duracion}</p>
      <p><strong>Cupo:</strong> {actividad.cupo}</p>

      <button onClick={handleInscribirse} className="btn-inscribirse">
        Inscribirse
      </button>
    </div>
  );
}

export default ActividadDetalle;
