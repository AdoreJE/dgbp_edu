package ieetubase.view;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
public class TestController {
	@RequestMapping("/index")
	public String index() throws Exception {
		return "index";
	}

}
