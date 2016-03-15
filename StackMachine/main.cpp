#include "stdafx.h"
#include <iostream>
#include <map>
#include <functional>
#include <exception>
#include "StackMachine.h"

using namespace std;


void execute(CStackMachine & stackMachine, FILE *input)
{
	int value = 0;
	char command[100];

	map<string, function<void()>> commands = {
		{ "push", bind(&CStackMachine::Push, &stackMachine, cref(value)) },
		{ "add", bind(&CStackMachine::Add, &stackMachine) },
		{ "mul", bind(&CStackMachine::Multiply, &stackMachine) },
		{ "sub", bind(&CStackMachine::Substract, &stackMachine) },
		{ "div", bind(&CStackMachine::Divide, &stackMachine) }
	};
	for (;;)
	{
		value = 0;
		cout << "input commands: ";
		int matchCount = fscanf(input, "%99s %d", command, &value);

		switch (matchCount)
		{
		case EOF:
			if (ferror(input))
			{
				throw std::runtime_error("input reading error");
			}
			return;
		case 0:
			throw std::logic_error("unexpected scan error");
		case 1:
		case 2:
			auto action = commands.at(command);
			action();
			stackMachine.PrintStack();
			break;
		}
	}
}

int main()
{
	CStackMachine stackMachine;

	stackMachine.Push(1);
	stackMachine.Push(2);
	stackMachine.PrintStack();

	try
	{
		execute(stackMachine, stdin);
		stackMachine.PrintStack();
	}
	catch (exception const& ex)
	{
		cout << ex.what() << endl;
		return 1;
	}

	return 0;
}

