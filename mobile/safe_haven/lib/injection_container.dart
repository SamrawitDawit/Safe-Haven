import 'package:get_it/get_it.dart';
import 'package:http/http.dart' as http;
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:safe_haven/features/auth/data/data_sources/local_data_source.dart';
import 'package:safe_haven/features/auth/data/data_sources/remote_data_source.dart';
import 'package:safe_haven/features/auth/data/repositories/authentication_repo_impl.dart';
import 'package:safe_haven/features/auth/domain/repositories/auth_repository.dart';
import 'package:safe_haven/features/auth/domain/usecases/forgot_password.dart';
import 'package:safe_haven/features/auth/domain/usecases/google_sign_in.dart';
import 'package:safe_haven/features/auth/domain/usecases/log_in.dart';
import 'package:safe_haven/features/auth/domain/usecases/log_out.dart';
import 'package:safe_haven/features/auth/domain/usecases/reset_password.dart';
import 'package:safe_haven/features/auth/domain/usecases/sign_up.dart';
import 'package:safe_haven/features/auth/presentation/bloc/bloc/auth_bloc_bloc.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'core/network/network_info.dart';

final sl = GetIt.instance;
Future<void> init() async {
  //! features - product
  // Bloc
  sl.registerFactory(() => AuthBlocBloc(sl(), sl(), sl(), sl(), sl(), sl()));

  // UseCases
  sl.registerLazySingleton(() => SignUpUsecase(authenticationRepository: sl()));
  sl.registerLazySingleton(() => LogInUsecase(authenticationRepository: sl()));
  sl.registerLazySingleton(() => LogoutUsecase(authenticationRepository: sl()));
  sl.registerLazySingleton(
      () => ForgotPasswordUsecase(authenticationRepository: sl()));
  sl.registerLazySingleton(
      () => ResetPasswordUseCase(authenticationRepository: sl()));

  sl.registerLazySingleton(
      () => GoogleSignInUseCase(authenticationRepository: sl()));

  //Repository
  sl.registerLazySingleton<AuthenticationRepository>(() =>
      AuthenticationRepoImpl(
          remoteDataSource: sl(), localDataSource: sl(), networkInfo: sl()));

  // Data sources
  sl.registerLazySingleton<AuthenticationRemoteDataSource>(
      () => AuthRemoteDataSourceImpl(client: sl()));

  sl.registerLazySingleton<AuthenticationLocalDataSource>(
      () => AuthLocalDataSourceImpl(sharedPreferences: sl()));

  //! core

  sl.registerLazySingleton<NetworkInfo>(() => NetworkInfoImpl(
      sl())); //unnamed requirement of networkinfoimpl: connectionchecker
  sl.registerLazySingleton(() => CustomHttpClient(
        client: sl(),
        authenticationLocalDataSource: sl(),
      ));

  //! External

  final sharedPreferences = await SharedPreferences.getInstance();
  sl.registerLazySingleton(() => sharedPreferences);
  sl.registerLazySingleton(() => http.Client());
  sl.registerLazySingleton(() => InternetConnectionChecker());
}
