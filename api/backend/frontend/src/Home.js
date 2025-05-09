import { useState, useEffect } from "react";
import "./Home.css";

export default function Home() {
  const [actividades, setActividades] = useState([]);
  const [busqueda, setBusqueda] = useState("");
  const [inscripciones, setInscripciones] = useState([]);
  const [detalle, setDetalle] = useState(null);

  useEffect(() => {
    fetch("http://localhost:8080/api/actividades")
      .then((res) => res.json())
      .then((data) => setActividades(data))
      .catch((err) => console.error("Error al obtener actividades:", err));
  }, []);

  const actividadesFiltradas = actividades.filter((actividad) => {
    const texto = `${actividad.categoria} ${actividad.fecha.dia} ${actividad.fecha.hora}`.toLowerCase();
    return texto.includes(busqueda.toLowerCase());
  });

  const inscribirse = (actividad) => {
    if (!inscripciones.includes(actividad.id_actividad)) {
      setInscripciones([...inscripciones, actividad.id_actividad]);
    }
  };

  const mostrarDetalle = (actividad) => {
    setDetalle(actividad);
  };

  return (
    <div className="home-container">
      <h1 className="title">Actividades Deportivas</h1>

      <input
        className="search-input"
        placeholder="Buscar por categoría, día u horario"
        value={busqueda}
        onChange={(e) => setBusqueda(e.target.value)}
      />

      {detalle && (
        <div className="detalle">
          <h2>{detalle.categoria}</h2>
          <p><strong>Día:</strong> {detalle.fecha.dia}</p>
          <p><strong>Horario:</strong> {detalle.fecha.hora} - {detalle.fecha.hora_fin}</p>
          <p><strong>Cupo disponible:</strong> {detalle.cupo ? "Sí" : "No"}</p>
          <p><strong>Estado:</strong> {detalle.estado ? "Activa" : "Inactiva"}</p>
          <button onClick={() => inscribirse(detalle)} disabled={!detalle.cupo} className="btn">
            {detalle.cupo ? "Inscribirse" : "Sin cupo"}
          </button>
        </div>
      )}

      <div className="actividades-grid">
        {actividadesFiltradas.map((actividad) => (
          <div key={actividad.id_actividad} className="card">
            <h3>{actividad.categoria}</h3>
            <p>{actividad.fecha.dia} - {actividad.fecha.hora}</p>
            <p>Cupo: {actividad.cupo ? "Disponible" : "Lleno"}</p>
            <div className="botones">
              <button onClick={() => mostrarDetalle(actividad)} className="btn secundario">Ver detalle</button>
              <button onClick={() => inscribirse(actividad)} disabled={!actividad.cupo} className="btn primario">
                {actividad.cupo ? "Inscribirse" : "Sin cupo"}
              </button>
            </div>
          </div>
        ))}
      </div>

      <h2 className="subtitle">Mis actividades</h2>
      <div className="actividades-grid">
        {actividades
          .filter((a) => inscripciones.includes(a.id_actividad))
          .map((actividad) => (
            <div key={actividad.id_actividad} className="card">
              <h3>{actividad.categoria}</h3>
              <p>{actividad.fecha.dia} - {actividad.fecha.hora}</p>
              <button onClick={() => mostrarDetalle(actividad)} className="btn secundario">Ver detalle</button>
            </div>
          ))}
      </div>
    </div>
  );
}
