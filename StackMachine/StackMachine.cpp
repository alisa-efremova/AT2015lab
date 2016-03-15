#include "stdafx.h"
#include "StackMachine.h"
#include <map>
#include <functional>
#include <exception>

using namespace std;

CStackMachine::CStackMachine()
{
}


CStackMachine::~CStackMachine()
{
}

void CStackMachine::Push(double value)
{
	m_stack.push_back(value);
}

double CStackMachine::Pop()
{
	//check
	double top = m_stack.back();
	m_stack.pop_back();
	return top;
}

void CStackMachine::Add()
{
	double top = Pop();
	m_stack.back() += top;
}

void CStackMachine::Multiply()
{
	double top = Pop();
	m_stack.back() *= top;
}

void CStackMachine::Substract()
{
	double top = Pop();
	m_stack.back() -= top;
}

void CStackMachine::Divide()
{
	double top = Pop();
	m_stack.back() /= top;
}

void CStackMachine::Dump()
{
	m_stack.clear();
}

void CStackMachine::PrintStack() const
{
	cout << "--- stack begin ---" << endl;
	for (auto & value : m_stack)
	{
		cout << value << endl;
	}
	cout << "--- stack end ---" << endl;
}