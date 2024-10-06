import 'package:equatable/equatable.dart';

class CaseEntity extends Equatable {
  final String? title;
  final String? description;
  final String? image_url;
  final String? submitter_id;
  final String? video_url;

  CaseEntity(
      {this.title,
      this.description,
      this.image_url,
      this.submitter_id,
      this.video_url});
  @override
  List<Object?> get props =>
      [title, description, image_url, video_url, submitter_id];
}
