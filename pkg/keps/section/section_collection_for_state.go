package section

import (
	"github.com/calebamiles/keps/pkg/keps/section/internal/rendering"
)

func ForProvisionalState(renderingInfo renderingInfoProvider) (Collection, error) {
	summary, err := New(rendering.SummaryName, renderingInfo)
	if err != nil {
		return nil, err
	}

	motivation, err := New(rendering.MotivationName, renderingInfo)
	if err != nil {
		return nil, err
	}

	infoWithSections := &renderingInfoAndSectionProvider{
		ss: []rendering.SectionProvider{
			summary,
			motivation,
		},
		renderingInfoProvider: renderingInfo,
	}

	readme, err := NewReadme(infoWithSections)
	if err != nil {
		return nil, err
	}

	c := &collection{
		sections: []section{
			summary,
			motivation,
			readme,
		},
	}

	c.readme = readme

	return c, nil
}

func ForImplementableState(renderingInfo renderingInfoProvider) (Collection, error) {
	summary, err := New(rendering.SummaryName, renderingInfo)
	if err != nil {
		return nil, err
	}

	motivation, err := New(rendering.MotivationName, renderingInfo)
	if err != nil {
		return nil, err
	}

	developerGuide, err := New(rendering.DeveloperGuideName, renderingInfo)
	if err != nil {
		return nil, err
	}

	operatorGuide, err := New(rendering.OperatorGuideName, renderingInfo)
	if err != nil {
		return nil, err
	}

	teacherGuide, err := New(rendering.TeacherGuideName, renderingInfo)
	if err != nil {
		return nil, err
	}

	graduationCriteria, err := New(rendering.GraduationCriteriaName, renderingInfo)
	if err != nil {
		return nil, err
	}

	infoWithSections := &renderingInfoAndSectionProvider{
		ss: []rendering.SectionProvider{
			summary,
			motivation,
			developerGuide,
			operatorGuide,
			teacherGuide,
			graduationCriteria,
		},
		renderingInfoProvider: renderingInfo,
	}

	readme, err := NewReadme(infoWithSections)
	if err != nil {
		return nil, err
	}

	c := &collection{
		sections: []section{
			summary,
			motivation,
			developerGuide,
			operatorGuide,
			teacherGuide,
			graduationCriteria,
			readme,
		},
	}

	c.readme = readme

	return c, nil
}
