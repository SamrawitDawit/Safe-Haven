import 'dart:convert';
import 'package:flutter_test/flutter_test.dart';
import 'package:safe_haven/features/auth/data/models/log_in_model.dart';
import 'package:safe_haven/features/auth/domain/entities/log_in_entity.dart';

import '../../../../Fixtures/dummy_data_reader.dart';

void main() {
  final testLogInModel = LogInModel(
      password: 'pass',
      fullName: 'user model',
      phoneNumber: '123',
      email: "",);

  test('should be a subclass of product model', () async {
    expect(testLogInModel, isA<LogInEntity>());
  });

  test('should return a JSON map with proper data', () async {
    //arrange

    final expectedJson = {
      'userType': 'normal',
      'password': 'pass',
      'fullName': 'user model',
      'phoneNumber': '123',
      'email': '',
      'anonymousDifferentiator': ''
    };

    //act
    final result = testLogInModel.toJson();

    //assert

    expect(result, expectedJson);
  });

  test('should return a valid log in model json', () async {
    //arrange

    final Map<String, dynamic> jsonData =
        json.decode(readJson('dummy_sign_up_response.json'));

    //act

    final result = LogInModel.fromJson(jsonData['data']);

    //assert
    print('exi');
    print(json);
    expect(result, testLogInModel);
  });

  test('logged in user model returns self', () async {
    final loggedInStuff = LoggedInModel(
        category: 'cat',
        language: 'lang',
        password: 'pass',
        fullName: 'user model',
        phoneNumber: '123',
        email: ''
        );

    final Map<String, dynamic> jsonData =
        json.decode(readJson('dummy_sign_up_response.json'));

    final result = LoggedInModel.fromjson(jsonData['data']);

    expect(result, loggedInStuff);
  });
}
