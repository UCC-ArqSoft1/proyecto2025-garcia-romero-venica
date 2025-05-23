import React, { useState } from 'react';
import './App.css';

function App() {
  const [usuario, setUsuario] = useState('');
  const [contrasena, setContrasena] = useState('');
  const [error, setError] = useState('');

  const handleLogin = (e) => {
    e.preventDefault();

    // Simulación de login básico
    if ((usuario === 'admin' && contrasena === 'admin123') || 
        (usuario === 'socio' && contrasena === 'socio123')) {
      setError('');
      alert('Login correcto'); // redirección futura
    } else {
      setError('Usuario o contraseña incorrectos.');
    }
  };

  return (
    <div className="login-container">
      <form className="login-form" onSubmit={handleLogin}>
        <h2>Iniciar sesión</h2>
        <input
          type="text"
          placeholder="Usuario"
          value={usuario}
          onChange={(e) => setUsuario(e.target.value)}
        />
        <input
          type="password"
          placeholder="Contraseña"
          value={contrasena}
          onChange={(e) => setContrasena(e.target.value)}
        />
        <button type="submit">Ingresar</button>

        {/* Mensaje de error si hay */}
        {error && <div className="error-message">{error}</div>}
      </form>
    </div>
  );
}

export default App;
