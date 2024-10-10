import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:safe_haven/features/auth/domain/entities/log_in_entity.dart';
import 'package:safe_haven/features/auth/domain/usecases/log_in.dart';

import '../../../helpers/test_helper.mocks.dart';

void main() {
  late LogInUsecase logInUsecase;
  late MockAuthenticationRepository mockAuthenticationRepository;

  setUp(() {
    mockAuthenticationRepository = MockAuthenticationRepository();
    logInUsecase =
        LogInUsecase(authenticationRepository: mockAuthenticationRepository);
  });

  final testloginEnity =
      LogInEntity(password: 'password', fullName: '123', email: 'email');

  test(
      'tests the correct url is called when logging in and returns the proper logged in entity',
      () async {
    //arrange
    when(mockAuthenticationRepository.logIn(testloginEnity))
        .thenAnswer((_) async => Right(unit));
    //act
    final result = await logInUsecase(LogInParams(loginEntity: testloginEnity));
    //assert
    expect(result, Right(unit));
  });
}
