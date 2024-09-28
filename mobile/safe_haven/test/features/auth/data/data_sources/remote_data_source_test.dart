import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:safe_haven/features/auth/data/data_sources/remote_data_source.dart';
import 'package:safe_haven/features/auth/data/models/authenticated_model.dart';
import 'package:safe_haven/features/auth/data/models/log_in_model.dart';
import 'package:safe_haven/features/auth/data/models/sign_up_model.dart';
import 'package:http/http.dart' as http;

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late MockHttpClient mockHttpClient;
  late AuthRemoteDataSourceImpl authRemoteDataSourceImpl;

  setUp(() {
    mockHttpClient = MockHttpClient();
    authRemoteDataSourceImpl = AuthRemoteDataSourceImpl(client: mockHttpClient);
  });

  final tLoginModel =
      LogInModel(userType: 'normal', email: 'email', password: 'password');
  final tRegisterModel = SignUpModel(
    fullName: 'name',
      userType: 'normal',
      category: 'cat',
      language: 'lang',
      password: 'password',
      email: '123');
  group('login test', () {
    test('should return authenticated model(token) when successful', () async {
      final testloginresponse = jsonEncode({
        'statusCode': 201,
        'message': 'hi',
        'data': {
          'access_token':
              'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZ21haWwuY29tIiwic3ViIjoiNjZiZGUzNmU5YmJlMDdmYzM5MDM0Y2RkIiwiaWF0IjoxNzI0MTQ0MjQzLCJleHAiOjE3MjQ1NzYyNDN9.oyC9gsD5ozRSCRMsC8M5WE8Wwxyzsbcn6-l7dLS8fsQ'
        }
      });
      final authModel = AuthenticatedModel(
        refreshToken: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZ21haWwuY29tIiwic3ViIjoiNjZiZGUzNmU5YmJlMDdmYzM5MDM0Y2RkIiwiaWF0IjoxNzI0MTQ0MjQzLCJleHAiOjE3MjQ1NzYyNDN9.oyC9gsD5ozRSCRMsC8M5WE8Wwxyzsbcn6-l7dLS8fsQ',
          token:
              'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZ21haWwuY29tIiwic3ViIjoiNjZiZGUzNmU5YmJlMDdmYzM5MDM0Y2RkIiwiaWF0IjoxNzI0MTQ0MjQzLCJleHAiOjE3MjQ1NzYyNDN9.oyC9gsD5ozRSCRMsC8M5WE8Wwxyzsbcn6-l7dLS8fsQ');

      when(mockHttpClient.post(any, body: tLoginModel.toJson()))
          .thenAnswer((_) async => http.Response(testloginresponse, 200));

      final result = await authRemoteDataSourceImpl.login(tLoginModel);

      expect(result, authModel);
    });

    test('should return unit when signed up properly', () async {
      final tRegisterResponse = jsonEncode(
          {'fullName': 'lucky', 'password': 'pass', 'email': '@email'});
      when(mockHttpClient.post(any, body: tRegisterModel.toJson()))
          .thenAnswer((_) async => http.Response(tRegisterResponse, 201));

      final result = await authRemoteDataSourceImpl.signup(tRegisterModel);

      expect(result, unit);
    });
  });
}
