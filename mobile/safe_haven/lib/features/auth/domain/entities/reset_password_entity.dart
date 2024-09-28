import 'package:equatable/equatable.dart';

class ResetPasswordEntity extends Equatable {
  final String reset_token;
  final String new_password;

  ResetPasswordEntity({ required this.new_password , required this.reset_token});
  @override
  List<Object?> get props => [reset_token, new_password];
}
