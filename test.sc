fib(n){
	if(n == 0)
		return 0;
	if(n == 1)
		return 1;
	fib(n - 1) + fib(n - 2);
}
main(){
	print(fib(38));
}


