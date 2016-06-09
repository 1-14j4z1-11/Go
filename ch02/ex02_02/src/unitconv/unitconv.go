package unitconv

import (
	"fmt"
)

func WeightUnits(value float64) string {
	return fmt.Sprintf("[Weight]\t%f[kg]\t= %f[lb],\t%f[lb]\t= %f[kg]", value, value * 2.205, value, value * 0.4536)
}

func LengthUnits(value float64) string {
	return fmt.Sprintf("[Length]\t%f[m]\t= %f[yards],\t%f[yards]\t= %f[m]", value, value * 1.094, value, value * 0.9144)
}

func VelocityUnits(value float64) string {
	return fmt.Sprintf("[Velocity]\t%f[km/h]\t= %f[M/h],\t%f[M/h]\t= %f[km/h]", value, value * 0.6215, value, value * 1.609)
}

func ForceUnits(value float64) string {
	return fmt.Sprintf("[Force]\t\t%f[kgf]\t= %f[N],\t%f[N] =\t%f[kgf]", value, value * 9.80665, value, value * 0.10197)
}

func PressureUnits(value float64) string {
	return fmt.Sprintf("[Pressure]\t%f[Pa]\t= %f[mmHg],\t%f[mmHg]\t= %f[Pa]", value, value * 0.007502, value, value * 133.3)
}
