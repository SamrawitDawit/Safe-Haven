import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:http/http.dart' as http;
import 'package:safe_haven/core/constants/constants.dart';
import 'package:safe_haven/core/error/faliure.dart';
import 'package:safe_haven/features/auth/data/data_sources/local_data_source.dart';
import 'package:safe_haven/features/case/data/data_sources/local_data_source_case.dart';
import 'package:safe_haven/features/case/data/models/case_model.dart';

abstract class CaseRemoteDataSource {
  /// calls the http://createCase-endpoint
  Future<CaseModel> createCase(CaseModel caseModel);

  /// calls the http://updateCase-endpoint
  Future<CaseModel> updateCase(CaseModel caseModel);

  /// calls the http://deleteCase-endpoint
  Future<Unit> deleteCase(String id);

  /// calls the http://getCase-endpoint
  Future<CaseModel> getCase(CaseModel caseModel);

  /// calls the http://getCases-endpoint
  Future<List<CaseModel>> getCases();

  Future<CaseModel> createLoggedInCase(CaseModel caseModel, String submitterId);

  Future<CaseModel> createnotLoggedInCase(CaseModel caseModel);
}

class CaseRemoteDataSourceImpl implements CaseRemoteDataSource {
  final http.Client client;
  final AuthenticationLocalDataSource authenticationLocalDataSource;

  CaseRemoteDataSourceImpl(this.authenticationLocalDataSource,
      {required this.client});

  @override
  Future<CaseModel> createCase(CaseModel caseModel) async {
    Map<String, dynamic>? user = await authenticationLocalDataSource.getUser();
    print('this is where the user is called returned');
    print(user);

    if (user != null) {
      print('as logged in user is trying to report');
      String submitterId = user['id'];
      return createLoggedInCase(caseModel, submitterId);
    } else {
      print('still not logged in?');
      print('huh');
      print(user);
      print('this is the user returned');
      return createnotLoggedInCase(caseModel);
    }
  }

  @override
  Future<Unit> deleteCase(String id) async {
    var uri = Uri.parse('${Urls.caseUrl}/delete/:${id}');

    final response = await client.delete(uri);

    if (response.statusCode == 200) {
      return unit;
    } else {
      throw ServerException('error in deleting in data source');
    }
  }

  @override
  Future<CaseModel> getCase(CaseModel caseModel) {
    // var uri = Uri.parse('${Urls.caseUrl}/');
    throw UnimplementedError();
  }

  @override
  Future<List<CaseModel>> getCases() {
    throw UnimplementedError();
  }

  @override
  Future<CaseModel> updateCase(CaseModel caseModel) {
    // TODO: implement updateCase
    throw UnimplementedError();
  }

  @override
  Future<CaseModel> createLoggedInCase(
      CaseModel caseModel, String submitterId) async {
    var uri = Uri.parse('${Urls.caseUrl}/submit');

    try {
      print(jsonEncode(caseModel.toJson()));
      final response = await client.post(uri,
          body:
              jsonEncode(({...caseModel.toJson(), 'submitterId': submitterId})),
          headers: {
            'Content-Type': 'application/json',
          });
      print(response.body);
      if (response.statusCode == 201) {
        print(
          response.statusCode,
        );
        print('as not logged in user trying to report');
        print('ezi in the create case data source');
        final finalbody = response.body;
        return CaseModel.fromJson(json.decode(finalbody)['data']);
      } else {
        throw ServerException('error in remote data source in creating case');
      }
    } on ServerException {
      throw ServerException('server exception in remote data siurce case');
    }
  }

  @override
  Future<CaseModel> createnotLoggedInCase(CaseModel caseModel) async {
    // TODO: implement createnotLoggedInCase
    var uri = Uri.parse('${Urls.caseUrl}/submit');
    print('ezi in not logged intfrdfgyhuioiuhygtfdgtyuiop');
    // throw UnimplementedError();
    try {
      print(jsonEncode(caseModel.toJson()));
      final response = await client
          .post(uri, body: jsonEncode(caseModel.toJson()), headers: {
        'Content-Type': 'application/json',
      });
      print(response.body);
      print('thos os the freakn bidy');
      if (response.statusCode == 201) {
        print(
          response.statusCode,
        );
        print('ezi in not logged in');
        final finalbody = response.body;
        return CaseModel.fromJson(json.decode(finalbody)['data']);
      } else {
        throw ServerException('error in remote data source in creating case');
      }
    } on ServerException {
      throw ServerException('server exception in remote data siurce case');
    }
  }
}



// var request = http.MultipartRequest('POST', uri);
    // request.fields.addAll({
    //   'id': caseModel.id,
    //   'title': caseModel.title ?? '',
    //   'description': caseModel.description ?? '',
    //   'submitter_id': caseModel.submitter_id ?? '',
    //   'image_id': caseModel.image_url ?? ''
    // });
    // request.files.add(await http.MultipartFile.fromPath(
    //     'image', caseModel.image_url ?? '',
    //     contentType: MediaType('image', 'jpeg')));

    // var streamedResponse = await client.send(request);
    // final response = await http.Response.fromStream(streamedResponse);