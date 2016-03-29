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

int main(int argc, char * argv[])
{
	if (argc < 2)
	{
		cout << "This program calculates simple statements using stack machine." << endl
			<< "Execute program with path to the file (with commands for the stack machine), e.g. commands.txt" << endl;
		return 0;
	}
	
	FILE * inputFile;
	inputFile = fopen(argv[1], "r");

	if (!inputFile)
	{
		cout << "File name is incorrect" << endl;
		return -1;
	}

	CStackMachine stackMachine;

	try
	{
		execute(stackMachine, inputFile);
		stackMachine.PrintStack();
		fclose(inputFile);
	}
	catch (exception const& ex)
	{
		cout << ex.what() << endl;
		return 1;
	}

	return 0;
}

