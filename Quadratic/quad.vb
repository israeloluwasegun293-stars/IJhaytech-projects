Private Sub btnCompute_Click()

    ' Declare variables
    Dim a As Double
    Dim b As Double
    Dim c As Double
    Dim d As Double
    Dim x1 As Double
    Dim x2 As Double

    ' Get input values
    a = Val(txtA.Text)
    b = Val(txtB.Text)
    c = Val(txtC.Text)

    ' Check if quadratic
    If a = 0 Then
        txtResult.Text = "Not a quadratic equation"
        Exit Sub
    End If

    ' Calculate discriminant
    d = b ^ 2 - 4 * a * c

    If d > 0 Then
        x1 = (-b + Sqr(d)) / (2 * a)
        x2 = (-b - Sqr(d)) / (2 * a)

        txtResult.Text = "Two real roots: " & x1 & " and " & x2

    ElseIf d = 0 Then
        x1 = -b / (2 * a)

        txtResult.Text = "One real root: " & x1

    Else
        txtResult.Text = "No real roots (complex roots)"
    End If

End Sub