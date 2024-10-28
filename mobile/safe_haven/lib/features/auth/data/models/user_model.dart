import 'package:safe_haven/features/auth/domain/entities/user_entity.dart';

class UserModel extends UserEntity {
  UserModel(
      {required super.id,
      required super.fullname,
      required super.accessToken,
      required super.refreshToken,
      required super.resetToken,
      super.bio,
      required super.category,
      required super.language,
      required super.resetTokenExpiry,
      required super.createdAt,
      required super.updatedAt,
      required super.resetCode,
      super.imageUrl,
      super.phoneNumber,
      super.active,
      super.role,
      super.counselorAssigned,
      super.counselorId,
      super.googleSignin,
      super.lock,
      super.preferedContact,
      super.verified});

  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
        id: json['id'],
        fullname: json['fullname'],
        category: json['category'],
        language: json['language'],
        resetTokenExpiry: json['resetTokenExpiry'],
        createdAt: json['createdAt'],
        updatedAt: json['updatedAt'],
        resetCode: json['resetCode'],
        imageUrl: json['imageUrl'] ?? '',
        phoneNumber: json['phoneNumber'] ?? '',
        counselorAssigned: json['counselorAssigned'] ?? '',
        preferedContact: json['preferedContact'] ?? '',
        verified: json['verified'] ?? '',
        role: json['role'],
        active: json['active'],
        bio: json['bio'],
        accessToken: json['accessToken'],
        refreshToken: json['refreshToken'],
        resetToken: json['resetToken']);
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'fullname': fullname,
      'accessToken': accessToken,
      'refreshToken': refreshToken,
      'resetToken': resetToken,
      'bio': bio ?? '',
      'category': category,
      'language': language,
      'resetTokenExpiry': resetTokenExpiry,
      'createdAt': createdAt,
      'upsatedAt': updatedAt,
      'resetCode': resetCode,
      'imageUrl': imageUrl ?? '',
      'phoneNumber': phoneNumber ?? '',
      'active': active ?? '',
      'role': role ?? '',
      'counselorAssigned': counselorAssigned ?? '',
      'counselorId': counselorId ?? '',
      'googleSignIn': googleSignin ?? '',
      'lock': lock ?? '',
      'preferedContact': preferedContact ?? ''
    };
  }
}
