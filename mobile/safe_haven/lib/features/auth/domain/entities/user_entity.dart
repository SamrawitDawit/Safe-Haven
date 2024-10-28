import 'package:equatable/equatable.dart';

class UserEntity extends Equatable {
  final String id;
  final String fullname;
  final String accessToken;
  final String refreshToken;
  final String resetToken;
  final String? bio;
  final String? phoneNumber;
  final String? imageUrl;
  final String? role;
  final String? category;
  final String? active;
  final String? verified;
  final String? counselorAssigned;
  final String? preferedContact;
  final String? counselorId;
  final String language;
  final String resetTokenExpiry;
  final bool? googleSignin;
  final bool? lock;
  final String createdAt;
  final String updatedAt;
  final String resetCode;

  UserEntity(
      {required this.accessToken,
      required this.refreshToken,
      required this.resetToken,
      required this.id,
      required this.fullname,
      required this.bio,
      this.phoneNumber,
      this.imageUrl,
      this.role,
      required this.category,
      this.active,
      this.verified,
      this.counselorAssigned,
      this.preferedContact,
      this.counselorId,
      required this.language,
      required this.resetTokenExpiry,
      this.googleSignin,
      this.lock,
      required this.createdAt,
      required this.updatedAt,
      required this.resetCode});

  @override
  // TODO: implement props
  List<Object?> get props => throw UnimplementedError();
}
