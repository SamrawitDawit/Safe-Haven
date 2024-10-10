import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';
import 'package:safe_haven/core/error/faliure.dart';

abstract class BaseUsecase<ResultType, ParamsType> {
  Future<Either<Failure, ResultType>> call(ParamsType params);
}


class NoParams extends Equatable{
  @override
  List<Object?> get props => []; 
}