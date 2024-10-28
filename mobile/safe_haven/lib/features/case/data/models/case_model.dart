import 'dart:convert';

import 'package:safe_haven/features/case/domain/entities/case_entity.dart';

class CaseModel extends CaseEntity {
  CaseModel(
      {required super.id,
      super.location,
      super.title,
      super.description,
      super.image_url,
      super.submitter_id,
      super.video_url});

  factory CaseModel.fromJson(Map<String, dynamic> json) {
    return CaseModel(
        id: json['id'],
        title: json['title'],
        location: json['location'],
        description: json['description'],
        image_url: json['image_url'],
        video_url: json['video_url'],
        submitter_id: json['submitter_id']);
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'title': title ?? '',
      'description': description ?? '',
      'image_url': image_url ?? '',
      'video_url': video_url ?? '',
      'submitter_id': submitter_id,
    };
  }

  static toModel(CaseEntity caseEntity) {
    return CaseModel(
        id: 'id',
        title: caseEntity.title,
        description: caseEntity.description,
        submitter_id: caseEntity.submitter_id,
        image_url: caseEntity.image_url,
        video_url: caseEntity.video_url);
  }
}

String convertToJsonList(List<CaseModel> caseModels) {
  List<Map<String, dynamic>> jsonList =
      caseModels.map((singleCase) => singleCase.toJson()).toList();

  return json.encode(jsonList);
}
