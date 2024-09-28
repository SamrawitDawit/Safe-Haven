import 'package:safe_haven/core/error/exception.dart';
import 'package:safe_haven/features/auth/domain/entities/reset_password_entity.dart';

class ResetPasswordModel extends ResetPasswordEntity {
  ResetPasswordModel({required super.new_password, required super.reset_token});

  factory ResetPasswordModel.fromJson(Map<String, dynamic> json) {
    try {
      print('Parsing JSON: $json');
      return ResetPasswordModel(
          new_password: json['new_password'], reset_token: json['reset_token']);
    } catch (e) {
      throw JsonParsingException();
    }
  }

  Map<String, dynamic> toJson() {
    return {'token': reset_token, 'new_password': new_password};
  }

  static ResetPasswordModel toModel(ResetPasswordEntity resetPasswordEntity) {
    return ResetPasswordModel(
        new_password: resetPasswordEntity.new_password,
        reset_token: resetPasswordEntity.reset_token);
  }
}
