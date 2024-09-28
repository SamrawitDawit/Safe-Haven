import 'dart:convert';
import 'package:flutter_test/flutter_test.dart';
import 'package:safe_haven/features/auth/data/models/log_in_model.dart';
import 'package:safe_haven/features/auth/data/models/sign_up_model.dart';
import 'package:safe_haven/features/auth/domain/entities/sign_up_entity.dart';

import '../../../../Fixtures/dummy_data_reader.dart';

void main() {
  final testSignUpModel2 = SignUpModel(
      userType: 'normal',
      password: 'pass',
      fullName: 'user model',
      phoneNumber: '123',
      language: 'lang',
      category: 'cat',
      email: '',
      anonymousDifferentiator: '');

  test('should be a subclass of product model', () async {
    expect(testSignUpModel2, isA<SignUpEntity>());
  });

  test('should return a JSON map with proper data', () async {
    //arrange

    final expectedJson = {
      'userType': 'normal',
      'password': 'pass',
      'fullName': 'user model',
      'phoneNumber': '123',
      'anonymousDifferentiator': '',
      'email': '',
      'language': 'lang',
      'category': 'cat'
    };

    //act
    final result = testSignUpModel2.toJson();

    //assert

    expect(result, expectedJson);
  });

  test('should return a valid log in model json', () async {
    final testSignUpModel = SignUpModel(
      fullName: 'name',
        userType: 'anonymous',
        category: 'General',
        language: 'lang',
        password: 'pass');
    //arrange

    final Map<String, dynamic> jsonData =
        json.decode(readJson('dummy_sign_up_response.json'));

    //act

    final result = LogInModel.fromJson(jsonData['data']);

    //assert
    print('exi');
    print(json);
    expect(result, testSignUpModel);
  });
}
