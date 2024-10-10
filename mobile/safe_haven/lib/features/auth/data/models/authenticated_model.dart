import '../../../../core/error/exception.dart';
import '../../domain/entities/authenticated_entity.dart';

class AuthenticatedModel extends AuthenticatedEntity {
  AuthenticatedModel({required super.refreshToken, required super.token});

  factory AuthenticatedModel.fromJson(Map<String, dynamic> json) {
    try {
      print('Parsing JSON: $json');
      return AuthenticatedModel(
          token: json['accessToken'],
          refreshToken: json['refreshToken'],
          
          );
    } catch (e) {
      throw JsonParsingException();
    }
  }
}
