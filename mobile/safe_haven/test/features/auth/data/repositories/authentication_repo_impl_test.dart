import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:safe_haven/features/auth/data/models/authenticated_model.dart';
import 'package:safe_haven/features/auth/data/models/log_in_model.dart';
import 'package:safe_haven/features/auth/data/repositories/authentication_repo_impl.dart';
import 'package:safe_haven/features/auth/domain/entities/log_in_entity.dart';
import 'package:safe_haven/features/auth/domain/entities/sign_up_entity.dart';

import '../../../../helpers/test_helper.mocks.dart';

// checking if api is called properly

void main() {
  late AuthenticationRepoImpl authenticationRepoImpl;
  late MockAuthenticationRemoteDataSource mockAuthenticationRemoteDataSource;
  late MockAuthenticationLocalDataSource mockAuthenticationLocalDataSource;
  late MockNetworkInfo mockNetworkInfo;

  setUp(() {
    mockAuthenticationLocalDataSource = MockAuthenticationLocalDataSource();
    mockAuthenticationRemoteDataSource = MockAuthenticationRemoteDataSource();
    mockNetworkInfo = MockNetworkInfo();
    authenticationRepoImpl = AuthenticationRepoImpl(
        remoteDataSource: mockAuthenticationRemoteDataSource,
        localDataSource: mockAuthenticationLocalDataSource,
        networkInfo: mockNetworkInfo);
  });

  group('getconnectedsignup', () {
    final testsignup = SignUpEntity(
      fullName: 'name',
        language: 'lang',
        category: 'cat',
        password: 'pass',
        phoneNumber: '123');

    final testloginEnity =
        LogInEntity(userType: 'normal', password: 'password', phoneNumber: '123',fullName: 'name',email: '123');

    final testloginmodel =
        LogInModel(userType: 'normal', password: 'password', email: 'wer');

    test('should check device online', () async {
      when(mockNetworkInfo.isConnected).thenAnswer((_) async {
        return true;
      });
      when(mockAuthenticationRemoteDataSource.login(any)).thenAnswer((_) async {
        print('ezi');
        print(testloginmodel);
        return AuthenticatedModel(token: 'token' , refreshToken: 'refreshToken');
      });
      when(mockAuthenticationLocalDataSource.cacheToken('token'))
          .thenAnswer((_) async => unit);
      print('ezietach');
      print(testloginEnity);
      print(testloginEnity.email);
      print(testloginEnity.phoneNumber);
      print(testloginEnity.fullName);
      final result = await authenticationRepoImpl
          .logIn(LogInModel.toModel(testloginEnity));

      expect(result, const Right(unit));
    });
  });
}
