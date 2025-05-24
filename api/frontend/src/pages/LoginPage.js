import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./LoginPage.css";

function LoginPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Simulación temporal
    if (username === "admin" && password === "admin123") {
      navigate("/admin");
    } else if (username === "socio" && password === "socio123") {
      navigate("/socio");
    } else {
      setError("Usuario o contraseña incorrectos.");
    }

    // Conexión real (descomentar cuando esté listo el backend)
    /*
    try {
      const res = await fetch('http://TU_BACKEND/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password })
      });
      const data = await res.json();
      if (res.ok) {
        if (data.rol === 'admin') navigate('/admin');
        else if (data.rol === 'socio') navigate('/socio');
      } else {
        setError(data.message || 'Error de autenticación');
      }
    } catch (err) {
      setError('Error del servidor');
    }
    */
  };

  return (
    <div className="login-container">
      <h2>Iniciar Sesión</h2>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Usuario"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="Contraseña"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <button type="submit">Ingresar</button>
        {error && <div className="error-message">{error}</div>}
      </form>
    </div>
  );
}

export default LoginPage;

