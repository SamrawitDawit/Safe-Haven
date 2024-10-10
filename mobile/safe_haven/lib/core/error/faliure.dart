import 'package:equatable/equatable.dart';

abstract class Failure extends Equatable {
  final String errorMessage;

  const Failure(this.errorMessage);

  @override
  List<Object> get props => [errorMessage];
}

class ServerException extends Failure {
  const ServerException(super.errorMessage);
}

class SocketFailure extends Failure {
  const SocketFailure(super.errorMessage);
}

class ConnectionFailure extends Failure {
  const ConnectionFailure(super.errorMessage);
}

class DatabaseFailure extends Failure {
  const DatabaseFailure(super.errorMessage);
}

class CacheFailure extends Failure {
  const CacheFailure(super.errorMessage);
}

class NotFoundFailure extends Failure {
  const NotFoundFailure(super.errorMessage);
}

class UnauthorizedFailure extends Failure {
  const UnauthorizedFailure(super.errorMessage);
}
