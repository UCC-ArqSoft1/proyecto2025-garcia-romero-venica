package services

import (
	"backend/domain"
	"backend/dto"	
	usuarioClient "backend/clients/usuario"
	"errors"
)

func GetUsuarioDtoById(id int) (dto.UsuarioDto, error) {
	usuarioDomain, err := usuarioClient.GetUserById(id)
	if err != nil {
		return dto.UsuarioDto{}, err
	}

	return mapUsuarioToDto(usuarioDomain), nil
}

// Verificar credenciales y retornar el DTO si son válidas
func Login(email, password string) (dto.UsuarioDto, error) {
	usuarioDomain, valido := usuarioClient.VerifyCredentials(email, password)
	if !valido {
		return dto.UsuarioDto{}, errors.New("credenciales inválidas")
	}

	return mapUsuarioToDto(usuarioDomain), nil
}

// Función auxiliar para mapear de domain.Usuario a dto.UsuarioDto
func mapUsuarioToDto(u domain.Usuario) dto.UsuarioDto {
	return dto.UsuarioDto{
		ID:          u.ID,
		Nombre:      u.Nombre,
		Email:       u.Email,
		Password:    u.Password, // en un caso real esto no se debería exponer
		TipoUsuario: u.TipoUsuario,
	}
}