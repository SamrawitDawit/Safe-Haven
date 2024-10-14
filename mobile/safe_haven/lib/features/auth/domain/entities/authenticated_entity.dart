import 'package:equatable/equatable.dart';

class AuthenticatedEntity extends Equatable {
  final String token;
  final String refreshToken;

  AuthenticatedEntity( { required this.refreshToken, required this.token});
  @override
  List<Object?> get props => [token, refreshToken];
}
