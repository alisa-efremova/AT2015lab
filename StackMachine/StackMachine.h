#pragma once
#include "stdafx.h"
#include <iostream>
#include <vector>

static const int STACK_SIZE = 1000;

class CStackMachine
{
public:
	CStackMachine();
	~CStackMachine();

	void Push(double value);
	double Pop();
	void Add();
	void Multiply();
	void Substract();
	void Divide();
	void Dump();
	void PrintStack() const;
	void execute(FILE *input);

private:
	std::vector<double> m_stack;

};

