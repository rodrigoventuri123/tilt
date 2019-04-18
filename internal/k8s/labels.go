package k8s

import (
	"strconv"

	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/windmilleng/tilt/internal/model"
)

const TiltRunIDLabel = "tilt-runid"

var TiltRunID = uuid.New().String()

const ManifestNameLabel = "tilt-manifest"

const TiltDeployIDLabel = "tilt-deployid"

func TiltRunLabel() model.LabelPair {
	return model.LabelPair{
		Key:   TiltRunIDLabel,
		Value: TiltRunID,
	}
}

func TiltDeployLabel(dID model.DeployID) model.LabelPair {
	return model.LabelPair{
		Key:   TiltDeployIDLabel,
		Value: strconv.Itoa(int(dID)),
	}
}

func TiltRunSelector() labels.Selector {
	return labels.Set{TiltRunIDLabel: TiltRunID}.AsSelector()
}
