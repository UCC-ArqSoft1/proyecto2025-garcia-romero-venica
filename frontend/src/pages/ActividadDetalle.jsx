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
    fetch(`http://localhost:8080/actividades/${id}`, {
      headers: { Authorization: `Bearer ${token}` }
    })
      .then((res) => res.json())
      .then((data) => setActividad(data))
      .catch((err) => {
        console.error("Error al obtener actividad:", err);
        Swal.fire("Error", "No se pudo cargar la actividad", "error");
      });
  }, [id, token]);

  if (!actividad) {
    return <p className="detalle-container">Cargando actividad...</p>;
  }

  const imagenes = {
    futbol: "https://www.comunidad.madrid/sites/default/files/styles/image_style_16_9/public/img/colectivos/shutterstock_614527325.jpg?itok=znJHrFRa",
    yoga: "https://nutritionsource.hsph.harvard.edu/wp-content/uploads/2021/11/pexels-yan-krukov-8436601-copy-scaled.jpg",
    boxeo: "https://www.tagoya.com/blog/wp-content/uploads/2023/05/mejores-entrenamiento-saco-boxeo.jpg",
    tenis: "https://s3.abcstatics.com/media/bienestar/2019/08/02/tenis-abecedario-kgNF--1248x698@abc.jpg",
    natacion: "https://www.superprof.com.ar/blog/wp-content/uploads/2024/12/entrenamiento-natacion.jpg"
  };

  const imagen = imagenes[actividad.nombre.toLowerCase()] || "https://via.placeholder.com/800x400";

  const capitalizar = (texto) => texto.charAt(0).toUpperCase() + texto.slice(1);

  return (
    <main className="detalle-container">
      <button onClick={() => navigate(-1)} className="volver-btn">← Volver</button>
      <h1 className="detalle-titulo">{capitalizar(actividad.nombre)}</h1>
      <img src={imagen} alt={actividad.nombre} className="detalle-imagen" />
      <div className="detalle-info">
        <p><strong>Descripción:</strong> {actividad.descripcion}</p>
        <p><strong>Profesor:</strong> {actividad.profesor}</p>
        <p><strong>Horario:</strong> {actividad.horario}</p>
        <p><strong>Cupo total:</strong> {actividad.cupo}</p>
        <p><strong>Disponibles:</strong> {actividad.disponible}</p>
      </div>
    </main>
  );
}

export default ActividadDetalle;
