// SPDX-License-Identifier: MIT
package com.daimler.sechub.developertools.admin.ui.action.project;

import java.awt.event.ActionEvent;
import java.util.Optional;

import com.daimler.sechub.developertools.admin.ui.UIContext;
import com.daimler.sechub.developertools.admin.ui.action.AbstractUIAction;
import com.daimler.sechub.developertools.admin.ui.cache.InputCacheIdentifier;

public class FetchProjectFalsePositiveConfiguration extends AbstractUIAction {
	private static final long serialVersionUID = 1L;

	public FetchProjectFalsePositiveConfiguration(UIContext context) {
		super("Fetch project false positive configuration",context);
	}

	@Override
	public void execute(ActionEvent e) {
		Optional<String> projectId = getUserInput("Please enter project ID/name",InputCacheIdentifier.PROJECT_ID);
		if (! projectId.isPresent()) {
			return;
		}

		String data = getContext().getAdministration().fetchProjectFalsePositiveConfiguration(asSecHubId(projectId.get()));
		outputAsBeautifiedJSONOnSuccess(data);
	}

}