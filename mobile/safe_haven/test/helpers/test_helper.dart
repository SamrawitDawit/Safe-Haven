import 'package:http/http.dart' as http;
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:mockito/annotations.dart';
import 'package:safe_haven/core/network/network_info.dart';
import 'package:safe_haven/features/auth/data/data_sources/local_data_source.dart';
import 'package:safe_haven/features/auth/data/data_sources/remote_data_source.dart';
import 'package:safe_haven/features/auth/domain/repositories/auth_repository.dart';
import 'package:shared_preferences/shared_preferences.dart';

@GenerateMocks([
  AuthenticationRepository,
  AuthenticationRemoteDataSource,
  AuthenticationLocalDataSource,
  NetworkInfo,
  InternetConnectionChecker,
  SharedPreferences,
  CustomHttpClient,
], customMocks: [
  MockSpec<http.Client>(as: #MockHttpClient)
])
void main() {}
